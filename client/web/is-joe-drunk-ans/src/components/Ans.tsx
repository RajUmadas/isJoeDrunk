import * as React from 'react';
import {  RouteComponentProps } from 'react-router-dom';
import axios from 'axios';

interface IState {
    the_ans: any;
}

export default class Ans extends React.Component<RouteComponentProps, IState> {
    constructor(props: RouteComponentProps) {
        super(props);
        this.state = { the_ans: false }
    }
    public componentDidMount(): void {
        axios.get(`http://localhost:8000/`).then(data => {
            this.setState({ the_ans: data.data })
        })
    }

    public render() {
        const the_ans = this.state.the_ans;
        return (
            <div>
                    <div className="text-center">
                        <h2>{ the_ans['ans'] }</h2>
                    </div>
            </div>
        )
    }
}

