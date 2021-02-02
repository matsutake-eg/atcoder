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
        p:[usize;n],
    }

    let mut idx = vec![0; n + 1];
    for i in 0..n {
        idx[p[i]] = i;
    }

    let mut ans = Vec::new();
    let mut p = p;
    let mut dq = VecDeque::new();
    dq.push_back(1);
    while let Some(v) = dq.pop_front() {
        let s = idx[v];
        for i in (v..=s).rev() {
            let t = p[i - 1];
            p[i - 1] = p[i];
            p[i] = t;
            ans.push(i);
        }

        let mut mx = v;
        for i in (v - 1)..s {
            if p[i] != i + 1 {
                println!("{}", -1);
                return;
            }
            mx = max(mx, p[i]);
        }

        if s + 1 == n {
            if p[s] != s + 1 {
                println!("{}", -1);
                return;
            }
            break;
        }
        dq.push_back(mx + 1);
    }

    if ans.len() != n - 1 {
        println!("{}", -1);
        return;
    }

    for v in ans {
        println!("{}", v);
    }
}
