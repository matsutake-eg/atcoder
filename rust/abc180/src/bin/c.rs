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
    }

    let mut ans = HashSet::new();
    for i in 1..=n {
        if i * i > n {
            break;
        } else if n % i == 0 {
            ans.insert(i);
            ans.insert(n / i);
        }
    }
    let mut ans: Vec<usize> = ans.into_iter().collect();
    ans.sort();
    for v in ans {
        println!("{}", v);
    }
}
