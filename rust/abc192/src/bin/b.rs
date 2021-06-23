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
        s:Chars,
    }

    for (i, c) in s.into_iter().enumerate() {
        if i % 2 == 0 && c.is_uppercase() || i % 2 == 1 && c.is_lowercase() {
            println!("No");
            return;
        }
    }
    println!("Yes");
}
