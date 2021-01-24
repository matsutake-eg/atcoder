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
        s:[String;n],
    }

    let mut hs1 = HashSet::new();
    let mut hs2 = HashSet::new();
    for v in s {
        if v.starts_with("!") {
            hs2.insert(v);
        } else {
            hs1.insert(v);
        }
    }

    for v in hs1 {
        if hs2.contains(&format!("!{}", v)) {
            println!("{}", v);
            return;
        }
    }
    println!("satisfiable");
}
