package main

import pb "github.com/jakebjorke/shipper/vessel-service/proto/vessel"
import "gopkg.in/mgo.v2"
import "gopkg.in/mgo.v2/bson"

const (
	dbName           = "shipper"
	vesselCollection = "vessels"
)

//Repository is the repo interface requirements.
type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Create(vessel *pb.Vessel) error
	Close()
}

//VesselRepository is an implementation of the Repository interface for vessels
type VesselRepository struct {
	session *mgo.Session
}

//FindAvailable searches for available vessels
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	var vessel *pb.Vessel

	err := repo.collection().Find(bson.M{
		"capacity":  bson.M{"$gte": spec.Capacity},
		"maxweight": bson.M{"$gte": spec.MaxWeight},
	}).One(&vessel)
	if err != nil {
		return nil, err
	}

	return vessel, nil
}

//Create is used to generate and save a new vessel
func (repo *VesselRepository) Create(vessel *pb.Vessel) error {
	return repo.collection().Insert(vessel)
}

//Close closes the connection to the repository
func (repo *VesselRepository) Close() {
	repo.session.Close()
}

func (repo *VesselRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(vesselCollection)
}
