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
        s:[String;n],
    }

    println!("{}", dfs(&s));
}

fn dfs(s: &[String]) -> usize {
    let l = s.len();
    if l == 0 {
        return 1;
    }

    dfs(&s[..(l - 1)])
        + match s[l - 1].as_str() {
            "OR" => 1 << l,
            _ => 0,
        }
}
