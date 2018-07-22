package set_test

import (
	"strconv"
	"testing"

	"github.com/StudioSol/PalcoMP3/lib/set"
	. "github.com/smartystreets/goconvey/convey"
)

var giantStringSlice = make([]string, giangSliceLength)

func init() {
	for i := 0; i < giangSliceLength; i++ {
		giantStringSlice[i] = strconv.Itoa(i + 1)
	}
}

func TestLinkedHashSetStringAdd(t *testing.T) {
	Convey("Given LinkedHashSetString.Add", t, func() {
		Convey("It should not store elements that are already on the Set", func() {
			set := set.NewLinkedHashSetString()
			set.Add("0", "0")
			set.Add("0")
			So(set.Length(), ShouldEqual, 1)
		})
		Convey("It should store elements with the correct constraints", func() {
			set := set.NewLinkedHashSetString()
			set.Add("0", "1", "2", "99", "93", "32", "00", "01", "2")
			So(set.Length(), ShouldEqual, 8)
		})
	})
}

func TestLinkedHashSetStringRemove(t *testing.T) {
	Convey("Given LinkedHashSetString.Remove", t, func() {
		Convey("It should remove elements from a Set", func() {
			set := set.NewLinkedHashSetString()
			set.Add(giantStringSlice...)

			// first element
			set.Remove(giantStringSlice[0])
			set.Remove(giantStringSlice[0])
			set.Remove(giantStringSlice[0])
			set.Remove(giantStringSlice[0])
			// last element
			set.Remove(giantStringSlice[len(giantStringSlice)-1])
			// arbitrary elements
			set.Remove(giantStringSlice[1000], giantStringSlice[2000], giantStringSlice[3000])
			So(set.Length(), ShouldEqual, len(giantStringSlice)-5)
		})
	})
}

func TestLinkedHashSetStringIter(t *testing.T) {
	Convey("Given LinkedHashSetString.Iter", t, func() {
		Convey("It should iterate over all elements of the set respecting the insertion order", func() {
			set := set.NewLinkedHashSetString()
			set.Add(giantStringSlice...)

			var (
				i                  int
				somethingWentWrong bool
			)
			for value := range set.Iter() {
				if value != giantStringSlice[i] {
					somethingWentWrong = true
					break
				}
				i++
			}
			So(somethingWentWrong, ShouldBeFalse)
			So(i, ShouldEqual, giangSliceLength)
		})
	})
}

func TestLinkedHashSetStringLength(t *testing.T) {
	Convey("Given LinkedHashSetString.Length", t, func() {
		Convey("It should return the correct length of the Set", func() {
			set := set.NewLinkedHashSetString()
			set.Add("0", "1", "2", "99", "93", "32", "00", "01", "2")
			So(set.Length(), ShouldEqual, 8)
			set.Remove("1")
			So(set.Length(), ShouldEqual, 7)
			set.Remove("2", "99", "94")
			So(set.Length(), ShouldEqual, 5)
			set.Add("94")
			So(set.Length(), ShouldEqual, 6)
		})

		Convey("It should return the correct length of the Set no matter the length of the Set", func() {
			set := set.NewLinkedHashSetString()
			set.Add(giantStringSlice...)
			So(set.Length(), ShouldEqual, len(giantStringSlice))
		})
	})
}
