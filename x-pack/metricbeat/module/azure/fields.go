// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package azure

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "azure", asset.ModuleFieldsPri, AssetAzure); err != nil {
		panic(err)
	}
}

// AssetAzure returns asset data.
// This is the base64 encoded gzipped contents of module/azure.
func AssetAzure() string {
	return "eJzUls+OmzAQxu95ilGOkTYPwKFSpV566K135JgJnS7YaGbYij59RQCv+b9pe0h8iBTbfN/38wwWL/CKTQLmd814AFDSAhM43v4fDwAZimWqlLxL4NMBALq9UPqsLtpHGAs0gglcUM0B4EpYZJLctr6AMyW+y7dDmwoTyNnXVT+z4DGWiaWUSszZkAsrg+QrNr88Z9H8onA3vv9A+NyBoDLZBd3BkVF8zRZnhjHDB+wGHZAKLV0J46hT3BFyU+FoYZ14J8YQpX0c/BU0irVoPUX8D97hGObaAdjkcj4t2vrLT7Q6Weom061g0Za0NFVFLu/3H0/H+yC6tgkYt7Czpml/pTILXXN3mwYpECzQatQ3g5vUlyCRUvbvnrEgfP0yM8yoRCfk3bhOKzXaqc9Ha7ORefQqR+Fmwa0vq1oxfSvPp1H06TV2D9K18GZl8W+B3nNuIKRiTYGC+iwsIXBYX6Bzasghp+REjbP46HB9XhjybiAx5iTKzbMgDXk3kAT5jZ6nSH3c+YVm1FyMYGqs9bV78DdqSAt92hlO6R2p591PljWuDe+59Na3S3cly+Q4N45u9/juOcIdlHZ86wICo9bsMDuvX02ink3+4K3ehwyd8ScAAP//n2w3Eg=="
}
