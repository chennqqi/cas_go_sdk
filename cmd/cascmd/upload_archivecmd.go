// Copyright 2019 chennqqi
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//      http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/antihax/optional"
	"github.com/dustin/go-humanize"
	"github.com/google/subcommands"
	openapi "gogs.fastapi.org/gitadmin/cas/go"
	"gogs.fastapi.org/gitadmin/cas/treehash"
)

const (
	MEGABYTE                        = 1024 * 1024
	GIGABYTE                        = 1024 * MEGABYTE
	DEFAULT_NORMAL_UPLOAD_THRESHOLD = 100 * MEGABYTE       //   # 单文件的默认最大上传阈值
	RECOMMEND_MIN_PART_SIZE         = 16 * MEGABYTE        // # 推荐的最小分片大小
	MAX_PART_NUM                    = 10000                // # 最大的分片数目
	MAX_FILE_SIZE                   = 4 * 10000 * GIGABYTE // # 最大支持40TB的文件存储
)

func init() {
	subcommands.Register(&uploadArchiveCmd{}, "")
}

type uploadArchiveCmd struct {
	vaultName, localFile, desc string
	uploadId                   string
	partSize                   int64
}

func (*uploadArchiveCmd) Name() string     { return "upload_archive" }
func (*uploadArchiveCmd) Synopsis() string { return "upload a local file" }
func (*uploadArchiveCmd) Usage() string {
	return `upload_archive <params>:
  Upload a local file.
`
}

func (p *uploadArchiveCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.vaultName, "vault", "", `format cas://vault-name`)
	f.StringVar(&p.localFile, "local_file", "", "file to be uploaded")
	f.StringVar(&p.uploadId, "-upload_id", "", "file to be uploaded")
	f.StringVar(&p.desc, "-desc", "", "description of the file")
	f.Int64Var(&p.partSize, "-part-size", 0, "multipart upload part size")
}

func (p *uploadArchiveCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if p.localFile == "" {
		fmt.Println("local file is must parameter")
		return subcommands.ExitFailure
	}
	conf, err := loadConf("")
	if err != nil {
		fmt.Println("load conf error:", err)
		return subcommands.ExitFailure
	}

	p.vaultName = parseVaultName(p.vaultName)
	client := openapi.NewAPIClient(conf)
	archive := client.ArchiveApi

	st, err := os.Stat(p.localFile)
	if os.IsNotExist(err) || err != nil {
		fmt.Println("ERROR local file:", err)
		return subcommands.ExitFailure
	}
	size := st.Size()
	var desc = p.desc
	if desc == "" {
		desc = filepath.Base(p.localFile)
	}
	var uploadId = p.uploadId
	var nparts int64
	var partSize int64
	if (p.partSize != 0 && size >= DEFAULT_NORMAL_UPLOAD_THRESHOLD) ||
		p.partSize > 0 && p.partSize > RECOMMEND_MIN_PART_SIZE {
		var resum int64
		if p.partSize > 0 {
			if uploadId == "" {
				if p.partSize%RECOMMEND_MIN_PART_SIZE != 0 {
					fmt.Println("Error: partsize must be divided by 16MB!")
					return subcommands.ExitFailure
				}
				if p.partSize*MAX_PART_NUM < RECOMMEND_MIN_PART_SIZE {
					fmt.Println("specified partsize too small, will be adjusted")
				}
				if p.partSize > size {
					fmt.Println("specified partsize too large, will be adjusted")
					for p.partSize > size {
						p.partSize /= 2
					}
				}
			} else {
				fmt.Println("test_utils larger than 100MB, multipart upload will be used")
				p.partSize = 16 * 1024 * 1024
			}
			for p.partSize < size {
				p.partSize *= 2
			}
			nparts = size / p.partSize
			if size%p.partSize == 0 {
				nparts += 1
			}
			fmt.Printf("Use %s parts with partsize %s to upload\n",
				nparts, humanize.Bytes((uint64)(p.partSize)))

			var opt openapi.UIDVaultsVaultNameMultipartUploadsPostOpts
			if p.desc != "" {
				opt.XCasArchiveDescription = optional.NewString(p.desc)
			}

			//Initiate Multipart Upload
			resp, err := archive.UIDVaultsVaultNameMultipartUploadsPost(ctx,
				conf.AppId, p.vaultName, fmt.Sprintf("%d", p.partSize), &opt)
			if err != nil {
				fmt.Println("ERROR:", err)
				return subcommands.ExitFailure
			}
			uploadId = resp.Header.Get("x-cas-multipart-upload-id")
			//location = resp.Header.Get("Location")
		} else {
			var opt openapi.UIDVaultsVaultNameMultipartUploadsUploadIDGetOpts
			//TODO: args
			parts, _, err := archive.UIDVaultsVaultNameMultipartUploadsUploadIDGet(ctx,
				conf.AppId, p.vaultName, uploadId, &opt)
			if err != nil {
				fmt.Println("ERROR:", err)
				return subcommands.ExitFailure
			}
			partSize := parts.PartSizeInBytes
			//			parts.Parts
			nparts = size / partSize
			if size%partSize == 0 {
				nparts += 1
			}
			//这里有个问题, 如果已上传的不是连续的, 那么就会漏掉中间的分块
			if len(parts.Parts) > 0 {
				var resumebytes int64
				for i := 0; i < len(parts.Parts); i++ {
					part := &parts.Parts[i]
					values := strings.Split(part.RangeInBytes, "-")
					if len(values) > 1 {
						var v int64
						fmt.Sscanf("%d", values[1], &v)
						if v > resumebytes {
							resumebytes = v
						}
					}
				}
				if resumebytes == size {
					resumebytes = nparts
				} else {
					resum = resumebytes / partSize
				}
			}

			fmt.Printf("Resume last upload with partsize %d\n",
				humanize.Bytes(uint64(partSize)))
		}
		var start, end int64
		/*
			Part = namedtuple('Namespace', ['vault', 'upload_id', 'local_file',
			                'start', 'end', 'etag', 'tree_etag'])
		*/
		var etagList []string
		for i := 0; i < int(nparts); i++ {
			end = size
			if end > partSize {
				end = partSize
			}
			end -= 1
			//compute tree hash
			etag, treeEtag, err := treehash.ComputeHashFromFile(p.localFile,
				start, end-start+1, 0)
			if err != nil {
				//TODO:
				fmt.Println("ERROR:", err)
				return subcommands.ExitFailure
			}

			start += partSize
			etagList = append(etagList, treeEtag)
			if i < int(resum) {
				continue
			}
			fmt.Println("Uploading part %d", i+1)
			//upload part
			uploadcmd := &uploadPartCmd{
				vaultName: p.vaultName,
				uploadId:  uploadId,
				localFile: p.localFile,
				start:     start,
				end:       end,
				eTag:      etag,
				treeTag:   treeEtag,
			}
			uploadcmd.Execute(ctx, nil)
		}

		etree := treehash.ComputeHashFromList(etagList)
		completecmd := &completeMultiPartCmd{
			vaultName: p.vaultName,
			uploadId:  uploadId,
			size:      size,
			treeTag:   etree,
		}
		completecmd.Execute(ctx, nil)
		return subcommands.ExitSuccess
	}
	if size <= RECOMMEND_MIN_PART_SIZE && p.partSize > 0 {
		fmt.Println("test_utils smaller than 16MB, part-size will be ignored.")
		//post file
		postarchivecmd := postArchiveCmd{
			vaultName: p.vaultName,
			localFile: p.localFile,
		}
		postarchivecmd.Execute(ctx, nil)
	}

	fmt.Println()
	return subcommands.ExitSuccess
}
