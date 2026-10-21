const API_URL = 'http://localhost:8080'

async function request(path, options = {}) {
    const response = await fetch(API_URL + path, options)
    const text = await response.text()
    if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`)
    }
    return text ? JSON.parse(text) : null

}

export async function getArtikel(limit, offset, status) {
    return await request(`/artikel?limit=${limit}&offset=${offset}&status=${status}`)
}

export async function getArtikelByID(id) {
    return await request(`/artikel/${id}`)
}

export async function createArtikel(artikel) {
    return await request('/artikel', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(artikel),
    })
}

export async function updateArtikel(id, artikel) {
    return await request(`/artikel/${id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(artikel),
    })
}

export async function deleteArtikel(id) {
    return await request(`/artikel/${id}`, {
        method: 'DELETE',
    })
}
