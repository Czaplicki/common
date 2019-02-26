#pragma once

class Scene;

typedef unsigned int uint;

typedef void(*SystemPtr)(Scene& p_scene, int p_entity, float p_deltaTime);

#define SYSTEM(name) void name (Scene& p_scene, int p_entity, float p_deltaTime)
#define ENTITY(mask) (p_scene.GetEntityMask(p_entity) & mask) == mask

class Scene
{

	int m_maxEntities;
	int m_currentEntities = 0;
	int* m_mask;

	int m_componentCount;
	void** m_componentData;
	std::map<uint, int> m_componentIndex;

	SystemPtr m_updateSystems[64];
	SystemPtr m_drawSystems[32];

	int m_updateSystemCount = 0;
	int m_drawSystemCount = 0;

	uint m_CreateEntity();

public:

	Scene(int p_entityAmount);
	~Scene();

	template <typename T>
	void SubmitComponent();

	template <typename T>
	T& GetComponent(int p_entity);

	template <typename T>
	uint GetMask();

	void SubmitUpdateSystem(SystemPtr p_system);
	void SubmitDrawSystem(SystemPtr p_system);

	uint DestroyEntity(int p_entity);
	uint RegisterEntity(int p_entityMask);

	uint GetMaxEntities();
	uint GetCurrentEntities();
	uint GetEntityMask(int p_index);

	void Update(float p_deltaTime);
	void Draw(float p_deltaTime);

};

template<typename T>
inline void Scene::SubmitComponent()
{
	m_componentData[m_componentCount] = new T[m_maxEntities];

	m_componentIndex[typeid(T).hash_code()] = m_componentCount;

	++m_componentCount;
}

template<typename T>
inline T& Scene::GetComponent(int p_entity)
{
	return ((T*)m_componentData[m_componentIndex[typeid(T).hash_code()]])[p_entity];
}

template<typename T>
inline uint Scene::GetMask()
{
	return 1 << m_componentIndex[typeid(T).hash_code()];
}