package service

import service "github.com/IRFAN374/upSvc2/reposiotry/token"

type Middleware func(service.Repository) service.Repository
