

// collector

if (input.path == "mysql") {

    var dbname = input.data.name || 'rd'
    var QPS = 0
    var IPS = 0
    var COUNT = 0;

    try {
        var rs = db.query(dbname, "show  global  status like 'Question%';")
        if (rs && rs.length > 0) {
            QPS = parseInt(rs[0].Value)
        }
    } catch (e) {
        console.info(e.message)
    }

    try {
        var rs = db.query(dbname, "show global status like 'Com_commit';")
        if (rs && rs.length > 0) {
            IPS = parseInt(rs[0].Value)
        }
    } catch (e) {
        console.info(e.message)
    }

    try {
        var rs = db.query(dbname, "SHOW FULL PROCESSLIST;")
        if (rs) {
            COUNT = rs.length
        }
    } catch (e) {
        console.info(e.message)
    }

    var now = (new Date()).getTime();

    influx.write({
        name: "mysql",
        tags: {
            "name": dbname
        },
        fields: {
            "QPS": QPS,
            "IPS": IPS,
            "COUNT": COUNT
        },
        time: now
    })
}

if (input.path == "/geoip") {
    output.data = geoip.country("61.135.152.133")
}
