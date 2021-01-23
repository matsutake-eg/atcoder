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
        l:usize,
    }

    println!("{}", combinations(l - 1, 11));
}

fn combinations(n: usize, k: usize) -> usize {
    if k > n {
        return 1;
    }

    let (mut nu, mut de) = (1, 1);
    for i in 0..k {
        nu = nu * (n - i);
        de = de * (i + 1);
        if nu % de == 0 {
            nu /= de;
            de = 1;
        }
    }
    nu / de
}
