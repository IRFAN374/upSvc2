package service

import service "github.com/IRFAN374/upSvc2/reposiotry/user"

type Middleware func(service.Repository) service.Repository
