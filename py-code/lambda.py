
people = [
    {"name":"harry","house":"griffindor"},
    {"name":"Cho","house":"ravenclaw"},
    {"name":"draco","house":"slytherin"}

]



people.sort(key = lambda person : person["name"] )

print(people)
