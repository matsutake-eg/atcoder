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
        s:usize,
        d:usize,
        xy: [(usize,usize);n],
    }

    for (x, y) in xy {
        if x >= s || y <= d {
            continue;
        }
        println!("{}", "Yes");
        return;
    }
    println!("{}", "No");
}
