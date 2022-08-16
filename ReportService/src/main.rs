#![feature(proc_macro_hygiene, decl_macro)]

extern crate reqwest;
use reqwest::Error;

extern crate rocket;
use rocket::{get, routes};
use rocket_contrib::json::Json;

use serde::Deserialize;
use serde::Serialize;

use std::collections::HashMap;

use chrono::{DateTime, Utc};

#[derive(Deserialize, Serialize, Debug)]
struct OrderItemForReportDTO {
    pub id: u32,
    pub id_order: u32,
    pub id_article: u32,
    pub article_name: String,
    pub current_price: f32,
    pub quantity: u32,
    pub total_price: f32,
}

#[derive(Deserialize, Serialize, Debug)]
struct OrderForReportDTO {
    pub id_order: u32,
    pub restaurant_name: String,
    pub id_restaurant: u32,
    pub id_app_user: u32,
    pub id_employee: u32,
    pub id_deliverer: u32,
    pub order_status: String,
    pub total_price: f32,
    pub tip: f32,
    pub date_time: String,
    pub order_items_for_report_dto: Vec<OrderItemForReportDTO>,
}

#[derive(Deserialize, Serialize, Debug)]
struct ArticlePriceQuantity {
    pub total_price: f32,
    pub quantity: u32,
}

#[derive(Deserialize, Serialize, Debug)]
struct Report {
    pub map_restaurants: HashMap<&str, f32>,
    pub map_articles: HashMap<&str, ArticlePriceQuantity>,
    pub date: String,
}

fn api_call() -> Result<Vec<OrderForReportDTO>, Box<dyn std::error::Error>> {
    let orders: Vec<OrderForReportDTO> =
        reqwest::get("http://localhost:8084/api/orders/ordersForReport")?.json()?;
    return Ok(orders);
}

#[get("/getReports")] //Result<Json<Vec<OrderForReportDTO>>
fn get_reports() -> Result<Json<Report>, Error> {
    let mut orders: Vec<OrderForReportDTO> = api_call().ok().unwrap();

    let mut map_restaurants: HashMap<&str, f32> = HashMap::new();
    let mut map_articles: HashMap<&str, ArticlePriceQuantity> = HashMap::new();

    orders.iter_mut().for_each(|el| {
        *map_restaurants.entry(&el.restaurant_name).or_insert(0.0) += el.total_price;

        el.order_items_for_report_dto.iter_mut().for_each(|el2| {
            if map_articles.contains_key(&el2.article_name as &str) {
                map_articles.insert(&el2.article_name, ArticlePriceQuantity { 
                    total_price: map_articles.get(&el2.article_name as &str).unwrap().total_price + el2.total_price, 
                    quantity: map_articles.get(&el2.article_name as &str).unwrap().quantity + el2.quantity
                });
            } else {
                map_articles.insert(&el2.article_name, ArticlePriceQuantity { total_price: el2.total_price, quantity: el2.quantity });
            }
        });
    });

    println!("{:?}", map_restaurants);
    println!("{:?}", map_articles);
    let now: DateTime<Utc> = Utc::now();
    println!("Now: {}", now.format("%d.%m.%Y."));

    let mut report: Report = Report {
        map_restaurants: map_restaurants,
        map_articles: map_articles,
        date: (now.format("%d.%m.%Y.")).to_string()
    };

    Ok(Json(report))
}

fn main() {
    rocket::ignite()
        .mount("/api/reports", routes![get_reports])
        .launch();
}
