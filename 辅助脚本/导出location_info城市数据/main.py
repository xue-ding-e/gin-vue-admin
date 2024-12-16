import json
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker
from model import LocationInfo  # 导入您的数据模型

# 数据库连接设置
DATABASE_URI = 'postgresql://postgres:123456@localhost:5432/gva_插件开发_delete'

def get_data():
    engine = create_engine(DATABASE_URI)
    Session = sessionmaker(bind=engine)
    session = Session()

    # 查询所有省份
    provinces = session.query(LocationInfo.province).distinct().all()
    data = []

    for province_tuple in provinces:
        province = province_tuple[0]
        # 查询省份下的所有城市
        cities = session.query(LocationInfo.city).filter_by(province=province).distinct().all()
        city_list = []
        for city_tuple in cities:
            city = city_tuple[0]
            # 查询城市下的所有区县
            districts = session.query(LocationInfo.district).filter_by(province=province, city=city).distinct().all()
            district_list = []
            for district_tuple in districts:
                district = district_tuple[0]
                district_dict = {
                    "label": district,
                    "value": district
                }
                district_list.append(district_dict)
            city_dict = {
                "label": city,
                "value": city,
                "children": district_list
            }
            city_list.append(city_dict)
        province_dict = {
            "label": province,
            "value": province,
            "children": city_list
        }
        data.append(province_dict)

    session.close()
    return data

if __name__ == "__main__":
    data = get_data()
    with open('location-area-data-min.js', 'w', encoding='utf-8') as f:
        f.write('export const data=')
        # 设置 separators 参数，确保输出最小化
        json.dump(data, f, ensure_ascii=False, separators=(',', ':'))
        f.write(';')