import * as React from 'react';
import './App.css';
import { Switch, Route, withRouter, RouteComponentProps, Link } from 'react-router-dom';
import Ans from './components/Ans';

class App extends React.Component<RouteComponentProps<any>> {
  public render() {
    return (
      <div>
        <Switch>
          <Route path={'/'} exact component={Ans} />
        </Switch>
      </div>
    );
  }
}
export default withRouter(App);