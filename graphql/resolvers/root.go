package resolvers

type resolvers struct {
	*userResolver
}

func Init() *resolvers {
	return &resolvers{
		NewUserResolver(),
	}
}
