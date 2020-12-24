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
        x:[i64;n],
    }

    println!("{}", x.iter().fold(0, |acc, x| acc + x.abs()));
    println!("{}", (x.iter().fold(0, |acc, x| acc + x * x) as f64).sqrt());
    println!("{}", x.iter().map(|x| x.abs()).max().unwrap());
}
