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

    dfs("", n);
}

fn dfs(s: &str, n: usize) {
    if s.len() == n {
        let mut cnt = 0i32;
        for c in s.chars() {
            if c == '(' {
                cnt += 1;
            } else {
                cnt -= 1;
            }
            if cnt < 0 {
                return;
            }
        }
        if cnt == 0 {
            println!("{}", s);
        }
        return;
    }

    dfs(&[s, "("].concat(), n);
    dfs(&[s, ")"].concat(), n);
}
