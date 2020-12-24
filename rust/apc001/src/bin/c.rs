#![allow(unused_imports)]
use itertools::Itertools;
use itertools_num::ItertoolsNum as _;
use num_integer::*;
use petgraph::*;
use std::cmp::*;
use std::collections::*;
use std::f64::consts::*;
use superslice::Ext as _;

fn main() {
    let n = read_line::<usize>();

    let mut l = 0;
    println!("{}", l);
    let mut ls = read_line::<String>();
    if ls == "Vacant" {
        return;
    }

    let mut r = n - 1;
    println!("{}", r);
    let mut rs = read_line::<String>();
    if rs == "Vacant" {
        return;
    }

    loop {
        let m = (l + r) / 2;
        println!("{}", m);
        let ms = read_line::<String>();
        if ms == "Vacant" {
            return;
        }

        let v1 = (r - m) % 2;
        let v2 = (m - l) % 2;
        if (v1 == 1 && rs == ms) || (v1 == 0 && rs != ms) {
            l = m;
            ls = ms;
        } else if (v2 == 1 && ms == ls) || (v2 == 0 && ms != ls) {
            r = m;
            rs = ms;
        }
    }
}

fn get_line() -> String {
    let mut s = String::new();
    std::io::stdin().read_line(&mut s).ok();
    s.trim().to_string()
}

fn read_line<T>() -> T
where
    T: std::str::FromStr,
    <T as std::str::FromStr>::Err: std::fmt::Debug,
{
    get_line().parse().unwrap()
}
