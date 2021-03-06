var GameBox = React.createClass({
    getInitialState: function() {
        return { data: {} };
    },
    componentDidMount: function() {
        $.ajax({
            url: this.props.url,
            dataType: 'json',
            success: function(data) {
                this.setState({data: data});
            }.bind(this),
            error: function(xhr, status, err) {
                console.error(this.props.url, status, err.toString());
            }.bind(this)
        });
    },
    handleGameFormSubmit: function(game) {
        console.log(game);
        $.ajax({
            url: this.props.url,
            dataType: 'json',
            type: 'GET',
            data: {
                teamCode: game.teamCode,
                date: game.date
            },
            success: function(data) {
                console.log(data);
                this.setState({data: data});
            }.bind(this),
            error: function(xhr, status, err) {
                console.error(this.props.url, status, err.toString());
                this.setState({data: "undefined"});
            }.bind(this)
        });
    },
    render: function() {
        return (
            <div className="gameBox">
                <GameForm onGameSubmit={this.handleGameFormSubmit} />
                <GameBoxscore data={this.state.data}></GameBoxscore>
            </div>
        );
    }
});

var GameForm = React.createClass({
    handleSubmit: function(e) {
        e.preventDefault();
        var teamCode = this.refs.teamCode.getDOMNode().value.trim();
        var date = this.refs.date.getDOMNode().value.trim();
        if (!teamCode || !date) {
            return;
        }
        this.props.onGameSubmit({teamCode: teamCode, date: date});
        return;
    },
    render: function() {
        return (
            <div className="gameForm">
                <form className="commentForm" onSubmit={this.handleSubmit}>
                    <p>
                        <label>Team code: </label>
                        <input type="text" ref="teamCode" size="3" maxLength="3" />
                    </p>
                    <p>
                        <label>Date (yyyy-mm-dd): </label>
                        <input type="text" ref="date" size="10" maxLength="10" />
                    </p>
                    <p>
                        <input type="Submit" value="Get game" />
                    </p>
                </form>
            </div>
        );
    }
});

var GameBoxscore = React.createClass({
    render: function() {
        var data = this.props.data;
        if (typeof data.game !== 'undefined') {
            return (
                <div className="gameBoxscore">
                    <h2>{data.game.AwayTeamName} at {data.game.HomeTeamName}</h2>
                    <h3>{data.game.Venue} - {data.game.Status}</h3>
                    <p>{data.game.Gameday}</p>
                    <p>{data.linescore.HomeTeamRuns}</p>
                    <p>{data.linescore.AwayTeamRuns}</p>
                </div>
                );
        } else {
            return (
                <div>
                    Please enter a team code and a date on which the team played.
                </div>
                );
        }
        return (<div/>);
    }
});

React.render(
    <GameBox url="game" />,
    document.getElementById('content')
);
