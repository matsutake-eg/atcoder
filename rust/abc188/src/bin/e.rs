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
        m:usize,
        a:[i64;n],
        xy: [(Usize1, Usize1);m],
    }

    let mut a_min = a.iter().enumerate().collect::<Vec<(usize, &i64)>>();
    a_min.sort_by(|a, b| (&a.1).cmp(&b.1));

    let mut g = vec![vec![]; n];
    for (x, y) in xy {
        g[x].push(y);
    }

    let mut ans = std::i64::MIN;
    let mut seen = HashSet::new();
    for (idx, c) in a_min {
        let mut dq = VecDeque::new();
        dq.push_back(idx);
        while let Some(v) = dq.pop_front() {
            for &nx in g[v].iter() {
                if seen.contains(&nx) {
                    continue;
                }
                dq.push_back(nx);
                seen.insert(nx);
                ans = max(ans, a[nx] - *c);
            }
        }
    }
    println!("{}", ans);
}
