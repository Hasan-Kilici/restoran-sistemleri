use std::fs::{File, OpenOptions};
use std::io::{BufRead, BufReader, Write};
use std::path::Path;

fn main() {
    let mut input_line = String::new();
    println!(
        "Komut giriniz
  1 - Masaya Ürün Ekle
  2 - Mevcut Siparişleri Listele
  3 - Masayı Kontrol Et"
    );
    std::io::stdin()
        .read_line(&mut input_line)
        .expect("Failed to read line");
    let komut: i32 = input_line.trim().parse().expect("Input not an integer");
    match komut {
        1 => ekle(),
        2 => siparisler(),
        3 => kontrol(),
        _ => hata(),
    }
}

fn ekle() {
    let mut input_line = String::new();
    println!("Masa Numarası giriniz:");
    std::io::stdin()
        .read_line(&mut input_line)
        .expect("Failed to read line");
    let masaid: u32 = input_line.trim().parse().expect("Input not an integer");

    input_line.clear();
    println!("Yemek kimliği giriniz:");
    std::io::stdin()
        .read_line(&mut input_line)
        .expect("Failed to read line");
    let yemekid: u32 = input_line.trim().parse().expect("Input not an integer");

    let food_path = Path::new("./src/foods.csv");
    let food_file = match File::open(&food_path) {
        Ok(file) => file,
        Err(error) => {
            println!("Failed to open file: {}", error);
            return;
        }
    };
    let food_reader = BufReader::new(food_file);

    let mut yemek_adi = String::new();
    let mut fiyat: u32 = 0;

    for line in food_reader.lines().skip(1) {
        let line = match line {
            Ok(content) => content,
            Err(error) => {
                println!("Failed to read line: {}", error);
                return;
            }
        };
        let fields: Vec<&str> = line.split(',').collect();

        let id = match fields[0].parse::<u32>() {
            Ok(value) => value,
            Err(error) => {
                println!("Failed to parse id: {}", error);
                return;
            }
        };

        if id == yemekid {
            yemek_adi = fields[1].to_string();
            fiyat = match fields[2].parse::<u32>() {
                Ok(value) => value,
                Err(error) => {
                    println!("Failed to parse price: {}", error);
                    return;
                }
            };
            break;
        }
    }

    if yemek_adi.is_empty() {
        println!("Yemek id'si foods.csv dosyasında bulunamadı.");
        return;
    }

    let order_path = Path::new("./src/orders.csv");
    let mut order_file = match OpenOptions::new()
        .append(true)
        .create(true)
        .open(&order_path) {
        Ok(file) => file,
        Err(error) => {
            println!("Failed to open file: {}", error);
            return;
        }
    };
    let last_order_id = match BufReader::new(&order_file).lines().last() {
        Some(Ok(line)) => {
            match line.split(',').next().and_then(|s| s.parse().ok()) {
                Some(value) => value,
                None => {
                    println!("Failed to parse last order id");
                    return;
                }
            }
        }
        _ => 0
    };
    let order_id = last_order_id + 1;
    match writeln!(&mut order_file, "{},{},{},{}", order_id, masaid, yemek_adi, fiyat) {
        Ok(_) => {
            println!(
                "Masa {} için {} yemeği {} TL fiyatla eklendi.",
                masaid, yemek_adi, fiyat
            );
        }
        Err(error) => {
            println!("Failed to write to file: {}", error);
        }
    }
}

fn siparisler() {
    let mut orders = Vec::new();

    let file = File::open("./src/orders.csv").unwrap();
    let reader = BufReader::new(file);

    for line in reader.lines().skip(1) {
        let line = line.unwrap();
        let fields: Vec<&str> = line.split(',').collect();
        let order_name = fields[2].to_string();
        orders.push(order_name);
    }

    println!("{:?}", orders);
}

fn kontrol() {
    let mut orders = Vec::new();
    let mut total_price = 0;

    let mut input_line = String::new();
    println!("Masa Numarası giriniz");
    std::io::stdin()
        .read_line(&mut input_line)
        .expect("Failed to read line");
    let masano: u32 = input_line.trim().parse().expect("Input not an integer");

    let file = File::open("./src/orders.csv").unwrap();
    let reader = BufReader::new(file);

    for line in reader.lines().skip(1) {
        let line = line.unwrap();
        let fields: Vec<&str> = line.split(',').collect();
        let masaid = fields[1].parse::<u32>().unwrap();
        if masaid == masano {
            let order_name = fields[2].to_string();
            let order_price = fields[3].parse::<u32>().unwrap();
            orders.push((order_name, order_price));
            total_price += order_price;
        }
    }

    orders.push(("Toplam Fiyat".to_string(), total_price));
    println!("{:?}", orders);
}

fn hata() {
    println!("BÖYLE BIR KOMUT YOK!");
}
