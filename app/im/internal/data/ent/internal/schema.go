// Code generated by ent, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = `{"Schema":"beetle/app/im/internal/data/ent/schema","Package":"beetle/app/im/internal/data/ent","Schemas":[{"name":"ChatMessage","config":{"Table":""},"fields":[{"name":"created_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"immutable":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":0},"schema_type":{"mysql":"datetime"}},{"name":"deleted_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":1},"schema_type":{"mysql":"datetime"}},{"name":"from","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"comment":"发送者"},{"name":"sender","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"comment":"接收者"},{"name":"message","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":2147483647,"position":{"Index":2,"MixedIn":false,"MixinIndex":0},"comment":"消息内容"}],"hooks":[{"Index":0,"MixedIn":true,"MixinIndex":1}],"interceptors":[{"Index":0,"MixedIn":true,"MixinIndex":1}]},{"name":"LoadRecord","config":{"Table":""},"fields":[{"name":"created_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"immutable":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":0},"schema_type":{"mysql":"datetime"}},{"name":"user_id","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"comment":"用户ID"},{"name":"device_id","type":{"Type":4,"Ident":"uuid.UUID","PkgPath":"github.com/google/uuid","PkgName":"uuid","Nillable":false,"RType":{"Name":"UUID","Ident":"uuid.UUID","Kind":17,"PkgPath":"github.com/google/uuid","Methods":{"ClockSequence":{"In":[],"Out":[{"Name":"int","Ident":"int","Kind":2,"PkgPath":"","Methods":null}]},"Domain":{"In":[],"Out":[{"Name":"Domain","Ident":"uuid.Domain","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]},"ID":{"In":[],"Out":[{"Name":"uint32","Ident":"uint32","Kind":10,"PkgPath":"","Methods":null}]},"MarshalBinary":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"MarshalText":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"NodeID":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}]},"Scan":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"String":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"Time":{"In":[],"Out":[{"Name":"Time","Ident":"uuid.Time","Kind":6,"PkgPath":"github.com/google/uuid","Methods":null}]},"URN":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"UnmarshalBinary":{"In":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"UnmarshalText":{"In":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Value":{"In":[],"Out":[{"Name":"Value","Ident":"driver.Value","Kind":20,"PkgPath":"database/sql/driver","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Variant":{"In":[],"Out":[{"Name":"Variant","Ident":"uuid.Variant","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]},"Version":{"In":[],"Out":[{"Name":"Version","Ident":"uuid.Version","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]}}}},"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"comment":"设备号"}]}],"Features":["intercept","schema/snapshot"]}`
