struct Aeroplane {
    company: String,
    plane: String,
    years: u32,
    capacity: u32,
}

impl Aeroplane {
    fn aero_comp(&self) -> &str {
        return &self.company;
    }

    fn aero_plane(&self) -> &str {
        return &self.plane;
    }

    fn aero_years(&self) -> &u32 {
        return &self.years;
    }

    fn aero_capacity(&self) -> &u32 {
        return &self.capacity;
    }

    fn is_flyable(&self, used_yrs: &u32) -> bool {
        return &self.years < used_yrs;
    }

    fn is_full(&self, curr_size: &u32) -> bool {
        return &self.capacity >= curr_size;
    }
}
fn main() {
    let new_aero = Aeroplane {
        company: String::from("Airbus"),
        plane: String::from("A320"),
        years: 12,
        capacity: 200,
    };
    let aero_2 = Aeroplane {
        company: String::from("Boeing"),
        plane: String::from("737 max"),
        years: 20,
        capacity: 300,
    };

    let years: u32 = 30;
    let cap: u32 = 300;
    println!("{}", new_aero.aero_comp());
    println!("{}", new_aero.aero_plane());
    println!("{}", new_aero.aero_years());
    println!("{}", new_aero.aero_capacity());
    println!("{}", new_aero.aero_comp());
    println!("{}", new_aero.aero_years());
    println!("{}", new_aero.is_flyable(&years));
    println!("{}", new_aero.is_full(&cap));
    println!("{}", aero_2.is_full(&cap));
    println!("{}", aero_2.aero_plane());
}
