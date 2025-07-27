package core

type (
	App struct {
		BaseObject
		Host

		Services    []*Service    `json:"services,omitempty"`
		Databases   []*Database   `json:"databases,omitempty"`
		Caches      []*Cache      `json:"caches,omitempty"`
		Workerpools []*Workerpool `json:"workerpools,omitempty"`

		ShareInstance bool `json:"share_instance,omitempty"` // Whether to share instances across services
	}
)

func DefaultApp() *App {
	return &App{
		BaseObject: BaseObject{
			Name: "webapp",
		},
		ShareInstance: true,
		Services: []*Service{
			{
				BaseObject: BaseObject{
					Name: "web_service",
				},
				ExecutorType: ExecutorTypeGolang,
			},
		},
		Databases: []*Database{
			{
				BaseObject: BaseObject{
					Name: "web_database",
				},
				DatabaseType: DatabaseTypeSqlite,
			},
		},
	}
}

func (w *App) Deploy() error {
	return nil
}
