//provide full API for the 'Creative Refresh'

//define module
mod CreativeRefresh {
    //define struct
    pub struct Refresh {
        pub id: u32,
        pub message: String,
        pub status: bool,
    }

    //define constructor
    impl Refresh {
        pub fn new(id: u32, message: String, status: bool) -> Refresh {
            Refresh {
                id: id,
                message: message,
                status: status,
            }
        }
    }

    //define public method
    impl Refresh {
        pub fn update(&mut self, new_message: String) {
            self.message = new_message;
        }

        pub fn set_status(&mut self, new_status: bool) {
            self.status = new_status;
        }
    }

    //define static private method
    impl Refresh {
        fn generate_unique_id() -> u32 {
            return rand::thread_rng().gen_range(1000, 10000);
        }
    }

    //define public method with static private method
    pub fn create_refresh(message: String, status: bool) -> Refresh {
        let mut refresh = Refresh::new(Self::generate_unique_id(), message, status);
        refresh.update(message);
        refresh.set_status(status);
        return refresh;
    }

    //define public method to list all refresh
    pub fn list_all_refreshes() -> Vec<Refresh> {
        let refreshes: Vec<Refresh> = vec![
            Refresh::new(1000, "Hello World".to_string(), true),
            Refresh::new(1001, "Hi There".to_string(), false),
            Refresh::new(1002, "Creative Refresh".to_string(), true),
        ];
        return refreshes;
    }

}

fn main() {
    println!("Starting creative refresh!");
    let refresh = CreativeRefresh::create_refresh("Welcome to Creative Refresh".to_string(), true);
    println!("Created Refresh: {:?}", refresh);
    let refreshes = CreativeRefresh::list_all_refreshes();
    println!("Refreshes: {:?}", refreshes);
}