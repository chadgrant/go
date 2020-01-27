package metadata

var md *Metadata

func Get() *Metadata {
	if md == nil {
		md = &Metadata{
			Vendor:          Vendor,
			Group:           Group,
			Service:         Service,
			Friendly:        Friendly,
			Description:     Description,
			Repo:            Repo,
			BuildNumber:     BuildNumber,
			BuiltBy:         BuiltBy,
			BuildTime:       BuildTime,
			GitHash:         GitHash,
			GitBranch:       GitBranch,
			CompilerVersion: CompilerVersion,
			Version:         1,
		}
	}
	return md
}
