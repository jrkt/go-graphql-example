package resolvers

type RootResolver struct {}

type resolvers struct {
	*userResolver
}

func Init() *resolvers {
	return &resolvers{
		NewUserResolver(),
	}
}
