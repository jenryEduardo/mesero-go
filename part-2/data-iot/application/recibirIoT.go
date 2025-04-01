package application

import "lastapi/part-2/data-iot/domain"

type RecibirIoT struct {
	r domain.InterfaceIoT
}

func NewRecibirIoT (r domain.InterfaceIoT) *RecibirIoT {
	return &RecibirIoT{r:r}
}

func(c *RecibirIoT) Run()([]domain.IoT, error) {
	return c.r.Recibir()
}