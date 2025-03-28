function controlVideoByServer(url: string, userid: string) {
    return {
        timeupdate: (videoid: string, time: number) => {
            fetch(url, {
                method: "POST",
                headers: {
                    Accept: "application/json",
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    userID: userid,
                    videoID: videoid,
                    time: time,
                }),
            });
        },
        loadtimevideoServer: async (videoid: string) => {
            const result = await fetch(
                `${url}?userid=${userid}&videoid=${videoid}`
            );
            if (result && result.status === 200) {
                try {
                    const data = await result.json();
                    if (data && data.code === 200) {
                        localStorage.setItem(videoid, data.data);
                    }
                } catch (err) {
                    console.log("Error ", err)
                }
            }
        },
    };
}

export interface ControlVideo {
    timeupdate: (videoid: string, time: number) => void;
    loadtimevideoServer: (videoid: string) => Promise<void>;
}

export default controlVideoByServer;