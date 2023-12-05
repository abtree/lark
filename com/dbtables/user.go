package dbtables

type Users struct {
	ID       uint64 `grom:"primaryKey,not null,autoIncrement" json:"id,omitempty"`
	Account  string `grom:"uniqueIndex,not null,default" json:"account,omitempty"`
	Name     string `grom:"uniqueIndex,not null,default" json:"name,omitempty"`
	CreateAt int64  `grom:"not null,default" json:"createat,omitempty"`
	LoginAt  int64  `grom:"not null,default" json:"loginat,omitempty"`
	LogoutAt int64  `grom:"not null,default" json:"logoutat,omitempty"`
	Level    uint32 `grom:"not null,default" json:"level,omitempty"`
	Avatar   uint32 `grom:"not null,default" json:"avatar,omitempty"`
	Data     []byte `grom:"type:MEDIUMBLOB,not null,default" json:"data,omitempty"`
}
