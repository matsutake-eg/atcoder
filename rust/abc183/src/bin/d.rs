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
        w:i64,
        stp: [(usize,usize,i64);n],
    }

    let mut l = vec![0; 200_001];
    for (s, t, p) in stp {
        l[s] += p;
        l[t] += -p;
    }
    let l = l.iter().cumsum().collect::<Vec<i64>>();
    println!(
        "{}",
        if *l.iter().max().unwrap() > w {
            "No"
        } else {
            "Yes"
        }
    );
}
