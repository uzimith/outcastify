var React = require("react");
var _ = require('lodash');
var oboe = require("oboe");

var User = React.createClass({
  render() {
    var user = this.props.data;
    return(
        <tr>
          <td>{user.Name}</td>
          <td><input type="checkbox" name={`join[${user.Id}]`} checked/></td>
          <td><input type="text" name={`private[${user.Id}]`} value={`text${user.Id}`} /></td>
        </tr>
        );
  }
});

var UserList = React.createClass({
  load() {
    oboe(this.props.url)
      .done( data => this.setState({users: data}) )
      .fail( (error) => console.error(error) )
  },
  getInitialState() {
    return {
      users: []
    };
  },
  componentDidMount() {
    this.load();
  },
  render() {
    var userlist = this.state.users.map( (user, index) => <User data={user} key={index} />)
    return(
        <table>
          <tbody>
            <tr>
              <th>Name</th>
              <th>Join</th>
              <th>Secret</th>
            </tr>
            {userlist}
          </tbody>
        </table>
      );
  }
});

window.renderUserList = function(id) {
  React.render(
      <UserList url={`/user/list?room=${id}`} />,
      document.getElementById('userlist')
      );
}

