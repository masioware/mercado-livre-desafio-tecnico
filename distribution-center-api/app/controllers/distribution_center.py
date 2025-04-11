from flask import Blueprint, Flask

from app.services import distribution_center_service

distribution_center = Blueprint("Distribution Center", __name__)


@distribution_center.get("/distribuitioncenters")
def get_distribution_center_by_item():

    dc_list = distribution_center_service.get_distribution_centers()

    return {"distribuitionCenters": dc_list}


def start(app: Flask):
    app.register_blueprint(distribution_center)
