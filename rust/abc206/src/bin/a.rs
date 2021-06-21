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
    }

    let n = n * 108 / 100;
    if n < 206 {
        println!("Yay!");
    } else if n == 206 {
        println!("so-so");
    } else {
        println!(":(");
    }
}
