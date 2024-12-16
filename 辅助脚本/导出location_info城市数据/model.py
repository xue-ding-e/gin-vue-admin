from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy import Column, Integer, String, Float

Base = declarative_base()

class LocationInfo(Base):
    __tablename__ = 'location_info'
    id = Column(Integer, primary_key=True)
    original_id = Column(String)
    name = Column(String, nullable=False)
    type = Column(String)
    typecode = Column(String)
    business_type = Column(String)
    address = Column(String, nullable=False)
    location = Column(String)
    phone = Column(String)
    province = Column(String, nullable=False)
    city = Column(String, nullable=False)
    district = Column(String, nullable=False)
    major_category = Column(String)
    middle_category = Column(String)
    small_category = Column(String)