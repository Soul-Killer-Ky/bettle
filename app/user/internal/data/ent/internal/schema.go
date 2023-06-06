// Code generated by ent, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = `{"Schema":"beetle/app/user/internal/data/ent/schema","Package":"beetle/app/user/internal/data/ent","Schemas":[{"name":"Group","config":{"Table":""},"edges":[{"name":"users","type":"User","through":{"N":"group_members","T":"GroupMember"}},{"name":"created_by","type":"User","ref_name":"created_groups","unique":true,"inverse":true}],"fields":[{"name":"created_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"immutable":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":0},"schema_type":{"mysql":"datetime"}},{"name":"updated_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"update_default":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":1},"schema_type":{"mysql":"datetime"}},{"name":"deleted_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":2},"schema_type":{"mysql":"datetime"}},{"name":"type","type":{"Type":11,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"comment":"群类型"},{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":30,"validators":1,"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"comment":"名称"},{"name":"icon","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":50,"validators":1,"position":{"Index":2,"MixedIn":false,"MixinIndex":0},"comment":"图标"},{"name":"memo","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":50,"validators":1,"position":{"Index":3,"MixedIn":false,"MixinIndex":0},"comment":"简介"}],"hooks":[{"Index":0,"MixedIn":true,"MixinIndex":2}],"interceptors":[{"Index":0,"MixedIn":true,"MixinIndex":2}]},{"name":"GroupMember","config":{"Table":""},"edges":[{"name":"group","type":"Group","field":"group_id","unique":true,"required":true},{"name":"user","type":"User","field":"user_id","unique":true,"required":true}],"fields":[{"name":"created_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"immutable":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":0},"schema_type":{"mysql":"datetime"}},{"name":"deleted_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":1},"schema_type":{"mysql":"datetime"}},{"name":"group_id","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"user_id","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"role","type":{"Type":11,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":2,"MixedIn":false,"MixinIndex":0},"comment":"群身份"}],"hooks":[{"Index":0,"MixedIn":true,"MixinIndex":1}],"interceptors":[{"Index":0,"MixedIn":true,"MixinIndex":1}],"annotations":{"Fields":{"ID":["group_id","user_id"],"StructTag":null}}},{"name":"User","config":{"Table":""},"edges":[{"name":"friends","type":"User"},{"name":"joined_groups","type":"Group","ref_name":"users","through":{"N":"group_members","T":"GroupMember"},"inverse":true},{"name":"created_groups","type":"Group"}],"fields":[{"name":"created_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"immutable":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":0},"schema_type":{"mysql":"datetime"}},{"name":"updated_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"update_default":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":1},"schema_type":{"mysql":"datetime"}},{"name":"deleted_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":2},"schema_type":{"mysql":"datetime"}},{"name":"phone","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":20,"unique":true,"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"comment":"手机号"},{"name":"username","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":30,"unique":true,"nillable":true,"validators":1,"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"comment":"唯一用户标识"},{"name":"password","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":100,"validators":1,"position":{"Index":2,"MixedIn":false,"MixinIndex":0},"comment":"密码"},{"name":"nickname","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":20,"validators":1,"position":{"Index":3,"MixedIn":false,"MixinIndex":0},"comment":"昵称"},{"name":"avatar","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":50,"validators":1,"position":{"Index":4,"MixedIn":false,"MixinIndex":0},"comment":"头像"},{"name":"memo","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":50,"validators":1,"position":{"Index":5,"MixedIn":false,"MixinIndex":0},"comment":"个人简介"}],"hooks":[{"Index":0,"MixedIn":true,"MixinIndex":2}],"interceptors":[{"Index":0,"MixedIn":true,"MixinIndex":2}]}],"Features":["intercept","schema/snapshot"]}`
