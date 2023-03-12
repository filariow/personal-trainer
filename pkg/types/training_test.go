package types_test

import (
	"time"

	"github.com/filariow/personal-trainer/pkg/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Training", func() {

	var et types.Training
	ft := types.Training{
		Exercises: []types.Exercise{
			{Name: "test", DurationInSec: 5, Preparation: 5},
		},
	}

	Describe("Duration", func() {
		Context("with empty training", func() {
			It("should be 0", func() {
				Expect(et.Duration()).To(Equal(time.Second * 0))
			})
		})

		Context("with excercises", func() {
			It("should be non zero", func() {
				Expect(ft.Duration()).To(Equal(time.Second * 10))
			})
		})
	})

})
