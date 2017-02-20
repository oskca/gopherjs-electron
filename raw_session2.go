package electron

import "github.com/gopherjs/gopherjs/js"

// SessionModule version@1.4.15
//
// Manage browser sessions, cookies, cache, proxy settings, etc.
type SessionModule struct {
	*js.Object
	// A Session object, the default session object of the app.
	DefaultSession *Session `js:"defaultSession"`
	// If partition starts with persist:, the page will use a persistent session available to all pages in the app with the same partition. if there is no persist: prefix, the page will use an in-memory session. If the partition is empty then default session of the app will be returned. To create a Session with options, you have to ensure the Session with the partition has never been used before. There is no way to change the options of an existing Session object.
	FromPartition func(Partition string, Options *SessionModuleFromPartitionOptions) (Obj *Session) `js:"fromPartition"`
}

func GetSessionModule() *SessionModule {
	o := Get("session")
	return &SessionModule{
		Object: o,
	}
}

type SessionModuleFromPartitionOptions struct {
	*js.Object
	// Whether to enable cache.
	Cache bool `js:"cache"`
}
