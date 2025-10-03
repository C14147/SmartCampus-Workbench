package auth

import (
    "log"

    casbin "github.com/casbin/casbin/v2"
    "github.com/casbin/casbin/v2/model"
    fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
)

func NewEnforcer(modelPath, policyPath string) *casbin.Enforcer {
    m, err := model.NewModelFromFile(modelPath)
    if err != nil {
        log.Fatalf("failed to load model: %v", err)
    }
    a := fileadapter.NewAdapter(policyPath)
    e, err := casbin.NewEnforcer(m, a)
    if err != nil {
        log.Fatalf("failed to create enforcer: %v", err)
    }
    if err := e.LoadPolicy(); err != nil {
        log.Fatalf("failed to load policy: %v", err)
    }
    return e
}
