package service

import service "github.com/IRFAN374/upSvc2/token"

type Middleware func(service.Service) service.Service
