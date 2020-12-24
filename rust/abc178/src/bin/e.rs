#![allow(unused_imports)]
use itertools::Itertools;
use itertools_num::ItertoolsNum as _;
use num_integer::*;
use petgraph::*;
use proconio::fastout;
use proconio::input;
use proconio::marker::*;
use std::cmp::*;
use std::collections::*;
use std::f64::consts::*;

#[allow(unused)]
const MOD: usize = 1_000_000_007;

// CAUTION: fastout has bug that it is not work in 'match arm', so, in that case, you should remove it or change logic.
#[proconio::fastout]
fn main() {
    proconio::input! {
        n:usize,
        xy: [(i64,i64);n],
    }

    let s: Vec<i64> = xy.iter().map(|(x, y)| *x + *y).collect();
    let d: Vec<i64> = xy.iter().map(|(x, y)| *x - *y).collect();
    let (max_s, min_s) = (s.iter().max().unwrap(), s.iter().min().unwrap());
    let (max_d, min_d) = (d.iter().max().unwrap(), d.iter().min().unwrap());
    println!("{}", max(max_s - min_s, max_d - min_d));
}
