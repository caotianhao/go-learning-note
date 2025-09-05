package main

import (
	"go-cth/model"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetSeerElementRestraint1v1(t *testing.T) {
	Convey("Test GetSeerElementRestraint1v1", t, func() {
		So(GetSeerElementRestraint1v1(
			model.SeerElemMystery,
			model.SeerElemGround,
		), ShouldEqual, 0.5)
	})
}

func TestGetSeerElementRestraint1v2(t *testing.T) {
	Convey("Test GetSeerElementRestraint1v2", t, func() {
		So(GetSeerElementRestraint1v2(
			model.SeerElemWar,
			model.SeerElemFly, model.SeerElemSuper,
		), ShouldEqual, 0.75)
	})
}

func TestGetSeerElementRestraint2v1(t *testing.T) {
	Convey("Test GetSeerElementRestraint2v1", t, func() {
		So(GetSeerElementRestraint2v1(
			model.SeerElemSuper, model.SeerElemLight,
			model.SeerElemDark,
		), ShouldEqual, 1.5)
	})
}

func TestGetSeerElementRestraint2v2(t *testing.T) {
	Convey("Test GetSeerElementRestraint2v2", t, func() {
		So(GetSeerElementRestraint2v2(
			model.SeerElemSuper, model.SeerElemFire,
			model.SeerElemWar, model.SeerElemNature,
		), ShouldEqual, 1.375)
	})

	Convey("Test GetSeerElementRestraint2v2", t, func() {
		So(GetSeerElementRestraint2v2(
			model.SeerElemDimension, model.SeerElemDragon,
			model.SeerElemElectricity, model.SeerElemHoly,
		), ShouldEqual, 1.125)
	})
}
