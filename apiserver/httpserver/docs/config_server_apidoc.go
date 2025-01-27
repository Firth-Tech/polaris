/**
 * Tencent is pleased to support the open source community by making Polaris available.
 *
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software distributed
 * under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package docs

import (
	"github.com/emicklei/go-restful/v3"
	restfulspec "github.com/polarismesh/go-restful-openapi/v2"
	apiconfig "github.com/polarismesh/specification/source/go/api/v1/config_manage"
)

var (
	configConsoleApiTags = []string{"ConfigConsole"}
	configClientApiTags  = []string{"Client"}
)

func EnrichCreateConfigFileGroupApiDocs(r *restful.RouteBuilder) *restful.RouteBuilder {
	return r.
		Doc("创建配置文件组").
		Metadata(restfulspec.KeyOpenAPITags, configConsoleApiTags).
		Reads(apiconfig.ConfigFileGroup{}, "开启北极星服务端针对控制台接口鉴权开关后，需要添加下面的 header\nHeader "+
			" X-Polaris-Token: {访问凭据}\n ```\n{\n    \"name\":\"someGroup\",\n  "+
			"  \"namespace\":\"someNamespace\",\n    \"comment\":\"some comment\",\n  "+
			"  \"createBy\":\"ledou\"\n}\n```")
}

func EnrichQueryConfigFileGroupsApiDocs(r *restful.RouteBuilder) *restful.RouteBuilder {
	return r.
		Doc("搜索配置文件组").
		Metadata(restfulspec.KeyOpenAPITags, configConsoleApiTags).
		Param(restful.QueryParameter("namespace", "命名空间，不填表示全部命名空间").
			DataType(DataType_String).Required(false)).
		Param(restful.QueryParameter("group", "配置文件分组名，模糊搜索").
			DataType(DataType_String).Required(false)).
		Param(restful.QueryParameter("fileName", "配置文件名称，模糊搜索").
			DataType(DataType_String).Required(false)).
		Param(restful.QueryParameter("offset", "翻页偏移量 默认为 0").
			DataType(DataType_Integer).
			Required(false).DefaultValue("0")).
		Param(restful.QueryParameter("limit", "一页大小，最大为 100").
			DataType(DataType_Integer).
			Required(true).DefaultValue("100"))
}

func EnrichDeleteConfigFileGroupApiDocs(r *restful.RouteBuilder) *restful.RouteBuilder {
	return r.
		Doc("删除配置文件组").
		Metadata(restfulspec.KeyOpenAPITags, configConsoleApiTags).
		Param(restful.QueryParameter("namespace", "命名空间").DataType(DataType_String).Required(true)).
		Param(restful.QueryParameter("group", "配置文件分组").DataType(DataType_String).Required(true))
}

func EnrichUpdateConfigFileGroupApiDocs(r *restful.RouteBuilder) *restful.RouteBuilder {
	return r.
		Doc("更新配置文件组").
		Metadata(restfulspec.KeyOpenAPITags, configConsoleApiTags).
		Reads(apiconfig.ConfigFileGroup{}, "开启北极星服务端针对控制台接口鉴权开关后，需要添加下面的 header\nHeader "+
			" X-Polaris-Token: {访问凭据}\n ```\n{\n    \"name\":\"someGroup\",\n  "+
			"  \"namespace\":\"someNamespace\",\n    \"comment\":\"some comment\",\n  "+
			"   \"createBy\":\"ledou\"\n}\n```")
}

func EnrichCreateConfigFileApiDocs(r *restful.RouteBuilder) *restful.RouteBuilder {
	return r.
		Doc("创建配置文件").
		Metadata(restfulspec.KeyOpenAPITags, configConsoleApiTags).
		Reads(apiconfig.ConfigFile{}, "开启北极星服务端针对控制台接口鉴权开关后，需要添加下面的 header\nHeader "+
			" X-Polaris-Token: {访问凭据}\n ```{\n    \"name\":\"application.properties\",\n   "+
			" \"namespace\":\"someNamespace\",\n    \"group\":\"someGroup\",\n   "+
			" \"content\":\"redis.cache.age=10\",\n    \"comment\":\"第一个配置文件\",\n  "+
			"  \"tags\":[{\"key\":\"service\", \"value\":\"helloService\"}],\n  "+
			"  \"createBy\":\"ledou\",\n    \"format\":\"properties\"\n}\n```\n")
}

func EnrichGetConfigFileApiDocs(r *restful.RouteBuilder) *restful.RouteBuilder {
	return r.
		Doc("拉取配置").
		Metadata(restfulspec.KeyOpenAPITags, configConsoleApiTags).
		Param(restful.QueryParameter("namespace", "命名空间").DataType(DataType_String).Required(true)).
		Param(restful.QueryParameter("group", "配置文件分组").DataType(DataType_String).Required(true)).
		Param(restful.QueryParameter("name", "配置文件名").DataType(DataType_String).Required(true))
}

func EnrichQueryConfigFilesByGroupApiDocs(r *restful.RouteBuilder) *restful.RouteBuilder {
	return r.
		Doc("搜索配置文件").
		Metadata(restfulspec.KeyOpenAPITags, configConsoleApiTags).
		Param(restful.QueryParameter("namespace", "命名空间").DataType(DataType_String).Required(false)).
		Param(restful.QueryParameter("group", "配置文件分组").DataType(DataType_String).Required(false)).
		Param(restful.QueryParameter("offset", "翻页偏移量 默认为 0").DataType(DataType_Integer).
			Required(false).DefaultValue("0")).
		Param(restful.QueryParameter("limit", "一页大小，最大为 100").DataType(DataType_Integer).
			Required(true).DefaultValue("100"))
}

func EnrichSearchConfigFileApiDocs(r *restful.RouteBuilder) *restful.RouteBuilder {
	return r.
		Doc("搜索配置文件").
		Metadata(restfulspec.KeyOpenAPITags, configConsoleApiTags).
		Param(restful.QueryParameter("namespace", "命名空间").DataType(DataType_String).Required(false)).
		Param(restful.QueryParameter("group", "配置文件分组").DataType(DataType_String).Required(false)).
		Param(restful.QueryParameter("name", "配置文件").DataType(DataType_String).Required(false)).
		Param(restful.QueryParameter("tags", "格式：key1,value1,key2,value2").DataType(DataType_String).Required(false)).
		Param(restful.QueryParameter("offset", "翻页偏移量 默认为 0").DataType(DataType_Integer).
			Required(false).DefaultValue("0")).
		Param(restful.QueryParameter("limit", "一页大小，最大为 100").DataType(DataType_Integer).
			Required(true).DefaultValue("100"))
}

func EnrichUpdateConfigFileApiDocs(r *restful.RouteBuilder) *restful.RouteBuilder {
	return r.
		Doc("更新配置文件").
		Metadata(restfulspec.KeyOpenAPITags, configConsoleApiTags).
		Reads(apiconfig.ConfigFile{}, "开启北极星服务端针对控制台接口鉴权开关后，需要添加下面的 header\nHeader "+
			" X-Polaris-Token: {访问凭据}\n ```{\n    \"name\":\"application.properties\",\n   "+
			" \"namespace\":\"someNamespace\",\n    \"group\":\"someGroup\",\n   "+
			" \"content\":\"redis.cache.age=10\",\n    \"comment\":\"第一个配置文件\",\n   "+
			" \"tags\":[{\"key\":\"service\", \"value\":\"helloService\"}],\n   "+
			" \"createBy\":\"ledou\",\n    \"format\":\"properties\"\n}\n```\n")
}

func EnrichDeleteConfigFileApiDocs(r *restful.RouteBuilder) *restful.RouteBuilder {
	return r.
		Doc("删除配置文件").
		Metadata(restfulspec.KeyOpenAPITags, configConsoleApiTags).
		Param(restful.QueryParameter("namespace", "命名空间").DataType(DataType_String).Required(true)).
		Param(restful.QueryParameter("group", "配置文件分组").DataType(DataType_String).Required(true)).
		Param(restful.QueryParameter("name", "配置文件").DataType(DataType_String).Required(true)).
		Param(restful.QueryParameter("deleteBy", "操作人").DataType(DataType_String).Required(false))
}

func EnrichBatchDeleteConfigFileApiDocs(r *restful.RouteBuilder) *restful.RouteBuilder {
	return r.
		Doc("批量删除配置文件").
		Metadata(restfulspec.KeyOpenAPITags, configConsoleApiTags).
		Param(restful.QueryParameter("deleteBy", "操作人").DataType(DataType_String).Required(false)).
		Reads(apiconfig.ConfigFile{}, "开启北极星服务端针对控制台接口鉴权开关后，需要添加下面的 header\nHeader "+
			" X-Polaris-Token: {访问凭据}\n```[\n     {\n         \"name\":\"application.properties\",\n "+
			"        \"namespace\":\"someNamespace\",\n         \"group\":\"someGroup\"\n     }\n]\n```")
}

func EnrichExportConfigFileApiDocs(r *restful.RouteBuilder) *restful.RouteBuilder {
	return r.
		Doc("导出配置文件").
		Metadata(restfulspec.KeyOpenAPITags, configConsoleApiTags).
		Reads(apiconfig.ConfigFileExportRequest{}, "```[\n     {\n         \"namespace\":\"someNamespace\",\n "+
			"        \"groups\":[\"someGroups\"]\n     \"names\":[\"application.properties\"],\n         }\n]\n```")
}

func EnrichImportConfigFileApiDocs(r *restful.RouteBuilder) *restful.RouteBuilder {
	return r.
		Doc("导入配置文件").
		Metadata(restfulspec.KeyOpenAPITags, configConsoleApiTags).
		Param(restful.QueryParameter("namespace", "命名空间").DataType(DataType_String).Required(true)).
		Param(restful.QueryParameter("group", "配置文件分组").DataType(DataType_String).Required(false)).
		Param(restful.MultiPartFormParameter("conflict_handling",
			"配置文件冲突处理，跳过skip，覆盖overwrite").DataType(DataType_String).Required(true)).
		Param(restful.MultiPartFormParameter("config", "配置文件").DataType("file").Required(true)).
		Reads(apiconfig.ConfigFile{}, "开启北极星服务端针对控制台接口鉴权开关后，需要添加下面的 header\nHeader"+
			" X-Polaris-Token: {访问凭据}\n```[\n     {\n         \"name\":\"application.properties\",\n "+
			"       \"namespace\":\"someNamespace\",\n         \"group\":\"someGroup\"\n     }\n]\n```")
}

func EnrichPublishConfigFileApiDocs(r *restful.RouteBuilder) *restful.RouteBuilder {
	return r.
		Doc("发布配置文件").
		Metadata(restfulspec.KeyOpenAPITags, configConsoleApiTags).
		Reads(apiconfig.ConfigFileRelease{}, "开启北极星服务端针对控制台接口鉴权开关后，需要添加下面的 header\nHeader "+
			" X-Polaris-Token: {访问凭据}\n```{\n    \"name\":\"release-002\",\n   "+
			" \"fileName\":\"application.properties\",\n    \"namespace\":\"someNamespace\",\n   "+
			" \"group\":\"someGroup\",\n    \"comment\":\"发布第一个配置文件\",\n    \"createBy\":\"ledou\"\n}\n```")
}

func EnrichGetConfigFileReleaseApiDocs(r *restful.RouteBuilder) *restful.RouteBuilder {
	return r.
		Doc("获取配置文件最后一次全量发布信息").
		Metadata(restfulspec.KeyOpenAPITags, configConsoleApiTags).
		Param(restful.QueryParameter("namespace", "命名空间").DataType(DataType_String).Required(true)).
		Param(restful.QueryParameter("group", "配置文件分组").DataType(DataType_String).Required(true)).
		Param(restful.QueryParameter("name", "配置文件").DataType(DataType_String).Required(true))
}

func EnrichGetConfigFileReleaseHistoryApiDocs(r *restful.RouteBuilder) *restful.RouteBuilder {
	return r.
		Doc("获取配置文件发布历史记录").
		Metadata(restfulspec.KeyOpenAPITags, configConsoleApiTags).
		Param(restful.QueryParameter("namespace", "命名空间").DataType(DataType_String).Required(true)).
		Param(restful.QueryParameter("group", "配置文件分组").DataType(DataType_String).Required(false)).
		Param(restful.QueryParameter("name", "配置文件").DataType(DataType_String).Required(false)).
		Param(restful.QueryParameter("offset", "翻页偏移量 默认为 0").DataType(DataType_Integer).
			Required(false).DefaultValue("0")).
		Param(restful.QueryParameter("limit", "一页大小，最大为 100").DataType(DataType_Integer).
			Required(true).DefaultValue("100"))
}

func EnrichGetAllConfigFileTemplatesApiDocs(r *restful.RouteBuilder) *restful.RouteBuilder {
	return r.
		Doc("获取配置模板").
		Metadata(restfulspec.KeyOpenAPITags, configConsoleApiTags)
}

func EnrichCreateConfigFileTemplateApiDocs(r *restful.RouteBuilder) *restful.RouteBuilder {
	return r.
		Doc("创建配置模板").
		Metadata(restfulspec.KeyOpenAPITags, configConsoleApiTags)
}

func EnrichGetConfigFileForClientApiDocs(r *restful.RouteBuilder) *restful.RouteBuilder {
	return r.
		Doc("拉取配置").
		Metadata(restfulspec.KeyOpenAPITags, configClientApiTags).
		Param(restful.QueryParameter("namespace", "命名空间").DataType(DataType_String).Required(true)).
		Param(restful.QueryParameter("group", "配置文件分组").DataType(DataType_String).Required(true)).
		Param(restful.QueryParameter("fileName", "配置文件名").DataType(DataType_String).Required(true)).
		Param(restful.QueryParameter("version", "配置文件客户端版本号，刚启动时设置为 0").
			DataType(DataType_Integer).Required(true))
}

func EnrichWatchConfigFileForClientApiDocs(r *restful.RouteBuilder) *restful.RouteBuilder {
	return r.
		Doc("监听配置").
		Metadata(restfulspec.KeyOpenAPITags, configClientApiTags).
		Reads(apiconfig.ClientWatchConfigFileRequest{}, "通过 Http LongPolling 机制订阅配置变更。")
}
