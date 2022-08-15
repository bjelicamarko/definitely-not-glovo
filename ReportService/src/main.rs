#![feature(proc_macro_hygiene, decl_macro)]


extern crate reqwest;
use reqwest::Error;
use reqwest::ClientBuilder;

extern crate rocket;
use rocket::{get, routes};
use rocket::{*, http::Status};
use rocket_contrib::json::Json;
use rocket::response::status;

use serde::Deserialize;
use serde::Serialize;
use std::time::Duration;


use std::io::Read;

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

fn api_call() -> Result<Vec<OrderForReportDTO>, Box<dyn std::error::Error>> {
    let orders: Vec<OrderForReportDTO> = reqwest::get("http://localhost:8084/api/orders/ordersForReport")?.json()?;
    return Ok(orders);
}

#[get("/getReports")]
fn get_reports() -> Result<Json<Vec<OrderForReportDTO>>, Error> { 
    let temp: Vec<OrderForReportDTO> = api_call().ok().unwrap();
    Ok(Json(temp))
}

fn main() {
    rocket::ignite().mount("/api/reports", routes![get_reports]).launch();
}