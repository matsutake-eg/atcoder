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

fn main() {
    input! {
        n:usize,
    }

    let mut ans = 1;
    let mut sum = 1;
    loop {
        if sum >= n {
            println!("{}", ans);
            return;
        }
        ans += 1;
        sum += ans;
    }
}
