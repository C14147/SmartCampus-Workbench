import React, { useEffect, useState } from 'react'
import { ApiClient } from '../services/apiClient'

export const SchoolsPage: React.FC = () => {
  const [list, setList] = useState<any[]>([])

  useEffect(() => {
    ApiClient.get('/schools').then((data:any) => setList(data)).catch(()=>setList([]))
  }, [])

  return (
    <div>
      <h2>Schools</h2>
      <ul>
        {list.map(s => <li key={s.id}>{s.name} ({s.code})</li>)}
      </ul>
    </div>
  )
}
