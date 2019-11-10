export default function (patch, body) {
    const  headers = new Headers();
    // headers.append("Content-Type","multipart/form-data")
    return fetch(`/${patch}`,{
            method: 'POST',
            mode: 'cors',
            cache: 'default',
            headers,
            body
        })
}