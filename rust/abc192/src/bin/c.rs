#![allow(unused_imports)]
use itertools::Itertools;
use itertools_num::ItertoolsNum as _;
use num_integer::*;
use petgraph::*;
use proconio::{fastout, input, marker::*};
use std::cmp::*;
use std::collections::*;
use std::f64::consts::*;
use superslice::Ext as _;

// CAUTION: don't print in 'match arm' because of fastout's bug
#[fastout]
fn main() {
    input! {
        n:usize,
        k:usize,
    }

    let mut ans = n;
    for _ in 0..k {
        ans = solve(ans);
    }
    println!("{}", ans);
}

fn solve(x: usize) -> usize {
    let mut cmin = x.to_string().chars().collect::<Vec<_>>();
    cmin.sort();
    let cmax = cmin
        .iter()
        .rev()
        .collect::<String>()
        .parse::<usize>()
        .unwrap();
    let cmin = cmin
        .into_iter()
        .collect::<String>()
        .parse::<usize>()
        .unwrap();
    cmax - cmin
}
