package config

type Cache struct {
	CachePath string `config:"cache_path" toml:"cache_path" yaml:"cache_path" json:"cache_path"`
}

func (c *Cache) copy() *Cache {
	if c == nil {
		return nil
	}

	return &Cache{
		CachePath: c.CachePath,
	}
}
