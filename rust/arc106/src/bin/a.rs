#![allow(unused_imports)]
use itertools::Itertools;
use itertools_num::ItertoolsNum as _;
use num_integer::*;
use petgraph::*;
use proconio::marker::*;
use std::cmp::*;
use std::collections::*;
use std::f64::consts::*;

// CAUTION: fastout has bug that it is not work in 'match arm', so, in that case, you should remove it or change logic.
#[proconio::fastout]
fn main() {
    proconio::input! {
        n:usize,
    }

    for i in 1..=38 {
        for j in 1..=26 {
            if 3usize.pow(i as u32) + 5usize.pow(j as u32) == n {
                println!("{} {}", i, j);
                return;
            }
        }
    }
    println!("{}", -1);
}
