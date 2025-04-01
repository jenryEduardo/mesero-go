package application

import "lastapi/part-2/data-iot/domain"

type EnviarIoT struct {
	r domain.InterfaceIoT
}

func NewEnviar(r domain.InterfaceIoT) *EnviarIoT {
	return &EnviarIoT{r:r}
}

func (c *EnviarIoT) Run(iot domain.IoT) error {
	return c.r.Enviar(iot)
}