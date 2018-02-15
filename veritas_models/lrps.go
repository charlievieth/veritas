package veritas_models

import (
	"sort"

	"code.cloudfoundry.org/bbs/models"
)

type VeritasLRPS map[string]*VeritasLRP

func (l VeritasLRPS) Get(guid string) *VeritasLRP {
	lrp, ok := l[guid]
	if !ok {
		lrp = &VeritasLRP{
			ProcessGuid:            guid,
			ActualLRPGroupsByIndex: map[string]*models.ActualLRPGroup{},
		}
		l[guid] = lrp
	}
	return lrp
}

func (l VeritasLRPS) SortedByProcessGuid() []*VeritasLRP {
	lrps := []*VeritasLRP{}

	for _, lrp := range l {
		lrps = append(lrps, lrp)
	}

	sort.Sort(VeritasLRPSByProcessGuid(lrps))

	return lrps
}

type VeritasLRPSByProcessGuid []*VeritasLRP

func (a VeritasLRPSByProcessGuid) Len() int           { return len(a) }
func (a VeritasLRPSByProcessGuid) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a VeritasLRPSByProcessGuid) Less(i, j int) bool { return a[i].ProcessGuid < a[j].ProcessGuid }
