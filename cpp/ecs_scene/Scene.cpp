#include "stdafx.h"

#include "Scene.hpp"

Scene::Scene(int p_entityAmount)
	:m_maxEntities(p_entityAmount)
{
	m_componentData = new void*[32];
	m_mask = new int[m_maxEntities];

	for (int i = 0; i < m_maxEntities; ++i)
	{
		m_mask[i] = NULL;
	}
}
Scene::~Scene()
{
	delete[] m_mask;
	delete[] m_componentData;
}


void Scene::SubmitUpdateSystem(SystemPtr p_system)
{
	m_updateSystems[m_updateSystemCount] = p_system;
	++m_updateSystemCount;
}

void Scene::SubmitDrawSystem(SystemPtr p_system)
{
	m_drawSystems[m_drawSystemCount] = p_system;
	++m_drawSystemCount;
}


uint Scene::m_CreateEntity()
{
	uint entity;
	for (entity = 0; entity < m_maxEntities; ++entity)
	{
		if (m_mask[entity] == NULL)
		{
			return(entity);
		}
	}

	printf("Error!  No more entities left!\n");

	return(m_maxEntities);
}
uint Scene::DestroyEntity(int p_entity)
{
	if (m_mask[p_entity])
	{
		--m_currentEntities;
		m_mask[p_entity] = NULL;
	}

	return 0;
}
uint Scene::RegisterEntity(int p_entityMask)
{
	++m_currentEntities;

	uint entity = m_CreateEntity();

	m_mask[entity] = p_entityMask;

	return entity;
}

uint Scene::GetMaxEntities()
{
	return m_maxEntities;
}
uint Scene::GetCurrentEntities()
{
	return m_currentEntities;
}
uint Scene::GetEntityMask(int p_index)
{
	return m_mask[p_index];
}

void Scene::Update(float p_deltaTime)
{
	for (int entity = 0; entity < m_maxEntities; ++entity)
	{
		for (int system = 0; system < m_updateSystemCount; ++system)
		{
			m_updateSystems[system](*this, entity, p_deltaTime);
		}
	}
}

void Scene::Draw(float p_deltaTime)
{
	for (int entity = 0; entity < m_maxEntities; ++entity)
	{
		for (int system = 0; system < m_drawSystemCount; ++system)
		{
			m_drawSystems[system](*this, entity, p_deltaTime);
		}
	}
}